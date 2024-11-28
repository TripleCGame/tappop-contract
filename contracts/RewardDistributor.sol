// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ECDSA.sol";
import "./EIP712.sol";

interface IERC20 {
    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function balanceOf(address account) external view returns (uint256);
}

contract RewardDistributor is EIP712 {
    bool private initialized; // Whether it has been initialized
    address public admin; // Contract administrator address
    address public signer; // Signer address
    mapping(uint256 => bool) private isUniqueIdDistributed; // Whether the unique_id has been distributed
    mapping(string => bool) private isSeasonToDistributed; // Whether the season has been distributed season_to:true/false

    // Debug event
    event DebugLog(
        address tokenAddress,
        address to,
        uint256 balance,
        uint256 amount
    );

    constructor() {
        admin = msg.sender; // Set the deploying account as administrator
    }

    // Initialize: mainly initializes eip712
    function initialize() external onlyAdmin {
        require(!initialized, "initialize: Already initialized!");
        eip712Initialize("RewardDistributor", "1.0.0");
        initialized = true;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can call this function.");
        _;
    }

    // Verify signature
    function verifySignature(bytes32 hash, bytes calldata signature)
        internal
        view
        returns (bool)
    {
        return ECDSA.recover(hash, signature) == signer;
    }

    // Set signer
    function setSinger(address _signer) external onlyAdmin {
        signer = _signer;
    }

    // Define a struct to store token address and amount
    struct ReceiverInfo {
        address tokenAddress; // Token contract address
        uint256 amount; // Amount of tokens to distribute
    }

    function distributeHash(
        uint256 season,
        ReceiverInfo[] calldata receiverInfos,
        address to,
        uint256 uniqueId
    ) internal  returns (bytes32) {
        require(receiverInfos.length > 0, "invalid receiverInfos!");
        string memory receiverInfoStr = receiverInfosToString(receiverInfos);
        return
            _hashTypedDataV4(
                keccak256(
                    abi.encode(
                        keccak256(
                            "distribute(uint256 season,address to,uint256 uniqueId,string receiverInfos)"
                        ),
                        season,
                        to,
                        uniqueId,
                        keccak256(bytes(receiverInfoStr))
                    )
                )
            );
    }

    // Distribute tokens to specified recipients (assuming the administrator is the recipient)
    function distribute(
        uint256 season,
        ReceiverInfo[] calldata receiverInfos,
        address to,
        uint256 uniqueId,
        bytes calldata signature
    ) public  {
        string memory seasonTo = string(abi.encodePacked(season, to));
        require(!isUniqueIdDistributed[uniqueId], "UniqueId Already Distributed!");
        require(!isSeasonToDistributed[seasonTo], "Season To Already Distributed!");

        require(verifySignature(
            distributeHash(
                season,
                receiverInfos,
                to,
                uniqueId
        ), signature),"Invalid signature!");
        isUniqueIdDistributed[uniqueId] = true;
        isSeasonToDistributed[seasonTo] = true;

        // check all
        for (uint256 i = 0; i < receiverInfos.length; i++) {
            address tokenAddress = receiverInfos[i].tokenAddress;
            uint256 amount = receiverInfos[i].amount;

            require(tokenAddress != address(0), "Invalid token address.");
            require(amount > 0, "Invalid amount.");

            // Check balance
            uint256 balance = balanceOfToken(tokenAddress);

            // Output debug information
            emit DebugLog(tokenAddress, to, balance, amount);
            require(balance >= amount, "Insufficient balance.");
        }

        // send all
        for (uint256 i = 0; i < receiverInfos.length; i++) {
            address tokenAddress = receiverInfos[i].tokenAddress;
            uint256 amount = receiverInfos[i].amount;

            IERC20 token = IERC20(tokenAddress);
            bool sent = token.transfer(to, amount); // Adjust recipient address if needed
            require(
                sent,
                string(
                    abi.encodePacked(
                        "Token transfer failed for token index ",
                        uintToString(i)
                    )
                )
            );
        }
    }

    // Function to get the balance of an address for a specific token
    function balanceOfToken(address tokenAddress)
        internal
        view
        returns (uint256)
    {
        bytes memory data = abi.encodeWithSelector(
            bytes4(keccak256(bytes("balanceOf(address)"))),
            address(this)
        );
        (bool success, bytes memory result) = tokenAddress.staticcall(data);
        require(success, "External call failed");
        return abi.decode(result, (uint256));
    }

    // Function to convert ReceiverInfo array to string
    function receiverInfosToString(ReceiverInfo[] calldata receiverInfos) public pure returns (string memory) {
        string memory result = "";
        uint256 length = receiverInfos.length;
        
        for (uint256 i = 0; i < length; i++) {
            string memory tokenAddressStr = addressToString(receiverInfos[i].tokenAddress);
            string memory amountStr = uintToString(receiverInfos[i].amount);
            
            // Concatenate string representation of each struct
            result = string(abi.encodePacked(result, tokenAddressStr, amountStr));
        }
        return result;
    }

    // Helper function: convert address to string
    function addressToString(address _addr) internal pure returns (string memory) {
        bytes32 value = bytes32(uint256(uint160(_addr)));
        bytes memory alphabet = "0123456789abcdef";

        bytes memory str = new bytes(42);
        str[0] = '0';
        str[1] = 'x';
        for (uint256 i = 0; i < 20; i++) {
            str[2+i*2] = alphabet[uint8(value[i + 12] >> 4)];
            str[3+i*2] = alphabet[uint8(value[i + 12] & 0x0f)];
        }
        return string(str);
    }

    // Helper function: convert uint to string
    function uintToString(uint256 _i) internal pure returns (string memory) {
        if (_i == 0) {
            return "0";
        }
        uint256 j = _i;
        uint256 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint256 k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - (_i / 10) * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
