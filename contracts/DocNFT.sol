pragma solidity ^0.8.7;

import "./contracts/token/ERC721/ERC721.sol";
import "./contracts/utils/Counters.sol";
import "./contracts/access/Ownable.sol";
import "./contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract DocNFT is ERC721URIStorage, Ownable {

    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor() ERC721("Document NFT", "DOCNFT"){

    }

    function mintNFT(address recipient, string memory tokenURI) public onlyOwner returns(uint) {

        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;

    }
}


