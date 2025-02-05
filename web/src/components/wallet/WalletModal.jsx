import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, Button, VStack } from "@chakra-ui/react";
import { useWallet } from "../hooks/useWallet";

const WalletModal = ({ isOpen, onClose }) => {
  const { connectWallet } = useWallet();

  return (
    <Modal isOpen={isOpen} onClose={onClose} isCentered>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Select Wallet</ModalHeader>
        <ModalBody>
          <VStack spacing={4}>
            <Button w="full" colorScheme="orange" onClick={() => connectWallet("metamask")}>
              Metamask
            </Button>
            <Button w="full" colorScheme="teal" onClick={() => connectWallet("walletconnect")}>
              WalletConnect
            </Button>
          </VStack>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default WalletModal;
