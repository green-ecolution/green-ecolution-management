import { create } from "zustand";

type State = {
  isOpen: boolean;
};

type Actions = {
  toggle: () => void;
  open: () => void;
  close: () => void;
};

type Store = State & Actions;

const useSidePanelStore = create<Store>((set) => ({
  isOpen: false,
  toggle: () => set((state) => ({ isOpen: !state.isOpen })),
  open: () => set({ isOpen: true }),
  close: () => set({ isOpen: false }),
}));

export const useIsSidePanelOpen = () => useSidePanelStore((state) => state.isOpen);
export const useToggleSidePanel = () => useSidePanelStore((state) => state.toggle);

export default useSidePanelStore;
