import { create } from "zustand";

export const useWalletStore = create((set) => ({
  account: null,
  isConnected: false,
  setAccount: (account) => set({ account, isConnected: !!account }),
}));
