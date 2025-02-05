import { useState, useEffect } from "react";
import { ethers } from "ethers";

export const useWallet = () => {
  const [account, setAccount] = useState(null);
  const [isConnected, setIsConnected] = useState(false);

  const connectWallet = async (providerType) => {
    if (window.ethereum && providerType === "metamask") {
      const provider = new ethers.BrowserProvider(window.ethereum);
      try {
        const accounts = await provider.send("eth_requestAccounts", []);
        setAccount(accounts[0]);
        setIsConnected(true);
      } catch (error) {
        console.error("Connection failed:", error);
      }
    } else {
      alert("No wallet detected!");
    }
  };

  const disconnectWallet = () => {
    setAccount(null);
    setIsConnected(false);
  };

  useEffect(() => {
    if (window.ethereum) {
      window.ethereum.on("accountsChanged", (accounts) => {
        setAccount(accounts.length > 0 ? accounts[0] : null);
        setIsConnected(accounts.length > 0);
      });
    }
  }, []);

  return { account, isConnected, connectWallet, disconnectWallet };
};
