import { useState } from "react";
import { Button } from "@chakra-ui/react";
import WalletModal from "./WalletModal";
import { useWallet } from "../hooks/useWallet";

const WalletButton = () => {
  const { isConnected, account, disconnectWallet } = useWallet();
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      {isConnected ? (
        <Button colorScheme="blue" onClick={disconnectWallet}>
          {account.slice(0, 6)}...{account.slice(-4)} (Disconnect)
        </Button>
      ) : (
        <Button colorScheme="green" onClick={() => setIsOpen(true)}>
          Connect Wallet
        </Button>
      )}
      <WalletModal isOpen={isOpen} onClose={() => setIsOpen(false)} />
    </>
  );
};

export default WalletButton;
