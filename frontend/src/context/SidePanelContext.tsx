import { createContext } from "react";

export interface SidePanelContext {
  isCollapsed: boolean;
  setIsCollapsed: (isCollapsed: boolean) => void;
}

const SidePanelContext = createContext<SidePanelContext | null>(null);



export default SidePanelContext;
