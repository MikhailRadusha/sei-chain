// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import {IWasmd} from "./precompiles/IWasmd.sol";
import {IJson} from "./precompiles/IJson.sol";
import {IAddr} from "./precompiles/IAddr.sol";

contract CW20ERC20Wrapper is ERC20 {

    address constant WASMD_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001002;
    address constant JSON_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001003;
    address constant ADDR_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001004;

    string public Cw20Address;
    IWasmd public WasmdPrecompile;
    IJson public JsonPrecompile;
    IAddr public AddrPrecompile;

    constructor(string memory Cw20Address_, string memory name_, string memory symbol_) ERC20(name_, symbol_) {
        WasmdPrecompile = IWasmd(WASMD_PRECOMPILE_ADDRESS);
        JsonPrecompile = IJson(JSON_PRECOMPILE_ADDRESS);
        AddrPrecompile = IAddr(ADDR_PRECOMPILE_ADDRESS);
        Cw20Address = Cw20Address_;
    }

    // Queries
    function decimals() public view override returns (uint8) {
        string memory req = _curlyBrace(_formatPayload("tokenInfo", ""));
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        bytes memory decimalsBytes = JsonPrecompile.extractAsBytes(response, "decimals");
        return uint8(abi.decode(decimalsBytes, (uint8)));
    }

    function balanceOf(address owner) public view override returns (uint256) {
        require(owner != address(0), "ERC20: balance query for the zero address");
        string memory ownerAddr = _formatPayload("address", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
        string memory req = _curlyBrace(_formatPayload("balance", _curlyBrace(ownerAddr)));
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        return JsonPrecompile.extractAsUint256(response, "balance");
    }

    function reqBalanceOf(address owner) public view returns (string memory) {
        require(owner != address(0), "ERC20: balance query for the zero address");
        string memory ownerAddr = _formatPayload("address", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
        string memory req = _curlyBrace(_formatPayload("balance", _curlyBrace(ownerAddr)));
        return req;
    }

    function callWasmd(string memory req) public view returns (bytes memory) {
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        return response;
    }

    function callWasmdString(string memory req) public view returns (string memory) {
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        return abi.decode(response, (string));
    }

    // function bytesToUint256(bytes memory byteArray) public pure returns (uint256) {
    //     require(byteArray.length <= 32, "Array too long");

    //     uint256 value;
    //     for (uint i = 0; i < byteArray.length; i++) {
    //         value = (value << 8) | uint256(uint8(byteArray[i]));
    //     }
    //     return value;
    // }

    function getSeiAddr(address evmAddr) public view returns (string memory) {
        return AddrPrecompile.getSeiAddr(evmAddr);
    }

    // function getOwnerAddr(address owner) public view returns (string memory) {
    //     string memory ownerAddr = _formatPayload("owner", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
    //     return ownerAddr;
    // }

    // function req(address owner) public view returns (string memory) {
    //     string memory ownerAddr = _formatPayload("address", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
    //     string memory req = _curlyBrace(_formatPayload("balance", _curlyBrace(ownerAddr)));
    //     return req;
    // }

    // function response(address owner) public view returns (bytes memory) {
    //     require(owner != address(0), "ERC20: balance query for the zero address");
    //     string memory ownerAddr = _formatPayload("address", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
    //     string memory req = _curlyBrace(_formatPayload("balance", _curlyBrace(ownerAddr)));
    //     bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
    //     return response;
    // }

    function totalSupply() public view override returns (uint256) {
        string memory req = _curlyBrace(_formatPayload("token_info", "{}"));
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        return JsonPrecompile.extractAsUint256(response, "total_supply");
    }

    // function req() public view returns (string memory) {
    //     string memory req = _curlyBrace(_formatPayload("token_info", "{}"));
    //     return req;
    // }

    // function response() public view returns (bytes memory) {
    //     string memory req = _curlyBrace(_formatPayload("token_info", "{}"));
    //     bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
    //     return response;
    // }

    function allowance(address owner, address spender) public view override returns (uint256) {
        string memory o = _formatPayload("owner", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
        string memory s = _formatPayload("spender", _doubleQuotes(AddrPrecompile.getSeiAddr(spender)));
        string memory req = _curlyBrace(_formatPayload("allowance", _curlyBrace(_join(o, s, ","))));
        bytes memory response = WasmdPrecompile.query(Cw20Address, bytes(req));
        return JsonPrecompile.extractAsUint256(response, "allowance");
    }

    function allowanceReq(address owner, address spender) public view returns (string memory) {
        string memory o = _formatPayload("owner", _doubleQuotes(AddrPrecompile.getSeiAddr(owner)));
        string memory s = _formatPayload("spender", _doubleQuotes(AddrPrecompile.getSeiAddr(spender)));
        string memory req = _curlyBrace(_formatPayload("allowance", _curlyBrace(_join(o, s, ","))));
        return req;
    }

    // Transactions
    function approve(address spender, uint256 amount) public override returns (bool) {
        string memory spenderAddr = _formatPayload("spender", _doubleQuotes(AddrPrecompile.getSeiAddr(spender)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("increase_allowance", _curlyBrace(_join(spenderAddr, amt, ","))));
        _execute(bytes(req));
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function approveReq(address spender, uint256 amount) public returns (string memory) {
        string memory spenderAddr = _formatPayload("spender", _doubleQuotes(AddrPrecompile.getSeiAddr(spender)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("increase_allowance", _curlyBrace(_join(spenderAddr, amt, ","))));
        return req;
    }

    function executeReq(string memory req) public returns (bytes memory) {
        return _execute(bytes(req));
    }

    function transfer(address to, uint256 amount) public override returns (bool) {
        require(to != address(0), "ERC20: transfer to the zero address");
        string memory recipient = _formatPayload("recipient", _doubleQuotes(AddrPrecompile.getSeiAddr(to)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("transfer", _curlyBrace(_join(recipient, amt, ","))));
        _execute(bytes(req));
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function transferReq(address to, uint256 amount) public view returns (string memory) {
        require(to != address(0), "ERC20: transfer to the zero address");
        string memory recipient = _formatPayload("recipient", _doubleQuotes(AddrPrecompile.getSeiAddr(to)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("transfer", _curlyBrace(_join(recipient, amt, ","))));
        return req;
    }

    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        require(to != address(0), "ERC20: transfer to the zero address");
        string memory sender = _formatPayload("owner", _doubleQuotes(AddrPrecompile.getSeiAddr(from)));
        string memory recipient = _formatPayload("recipient", _doubleQuotes(AddrPrecompile.getSeiAddr(to)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("transfer_from", _curlyBrace(_join(_join(sender, recipient, ","), amt, ","))));
        _execute(bytes(req));
        emit Transfer(from, to, amount);
        return true;
    }

    function transferFromReq(address from, address to, uint256 amount) public view returns (string memory) {
        require(to != address(0), "ERC20: transfer to the zero address");
        string memory sender = _formatPayload("owner", _doubleQuotes(AddrPrecompile.getSeiAddr(from)));
        string memory recipient = _formatPayload("recipient", _doubleQuotes(AddrPrecompile.getSeiAddr(to)));
        string memory amt = _formatPayload("amount", _doubleQuotes(Strings.toString(amount)));
        string memory req = _curlyBrace(_formatPayload("transfer_from", _curlyBrace(_join(_join(sender, recipient, ","), amt, ","))));
        return req;
    }

    function _execute(bytes memory req) internal returns (bytes memory) {
        (bool success, bytes memory ret) = WASMD_PRECOMPILE_ADDRESS.delegatecall(
            abi.encodeWithSignature(
                "execute(string,bytes,bytes)",
                Cw20Address,
                bytes(req),
                bytes("[]")
            )
        );
        require(success, "CosmWasm execute failed");
        return ret;
    }

    function _formatPayload(string memory key, string memory value) internal pure returns (string memory) {
        return _join(_doubleQuotes(key), value, ":");
    }

    function _curlyBrace(string memory s) internal pure returns (string memory) {
        return string.concat("{", string.concat(s, "}"));
    }

    function _doubleQuotes(string memory s) internal pure returns (string memory) {
        return string.concat("\"", string.concat(s, "\""));
    }

    function _join(string memory a, string memory b, string memory separator) internal pure returns (string memory) {
        return string.concat(a, string.concat(separator, b));
    }
}