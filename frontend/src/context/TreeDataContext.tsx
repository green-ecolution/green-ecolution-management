import { createContext, useContext } from "react";

export interface Tree {
  lat: number;
  lng: number;
  name: string;
  id: number;
  neededWater?: number;
  status?: "healthy" | "neutral" | "unhealthy";
  treeType?: string;
  treeHeight?: number;
  treeDiameter?: number;
  treeAge?: number;
  treeCrown?: number;
  treeHealth?: number;
  treeCondition?: number;
  treeRootCondition?: number;
  treeRootDepth?: number;
  treeRootMoisture?: number;
  treeRootNutrition?: number;
  treeRootStructure?: number;
  treeRootFungi?: number;
  treeRootBacteria?: number;
  treeRootInsects?: number;
  treeRootAnimals?: number;
  treeRootWater?: number;
  image?: string;
}

export interface TreeDataContext {
  trees: Tree[];
}

export interface TreeDataContextProviderProps {
  children: React.ReactNode;
}

const TreeDataContext = createContext<TreeDataContext | null>(null);

export const TreeDataContextProvider = ({
  children,
}: TreeDataContextProviderProps) => {
  const treeMetaData = {
    treeType: "oak",
    treeHeight: 10,
    treeDiameter: 20,
    treeAge: 30,
    treeCrown: 40,
    treeHealth: 50,
    treeCondition: 60,
    treeRootCondition: 70,
    treeRootDepth: 80,
    treeRootMoisture: 90,
    treeRootNutrition: 100,
    treeRootStructure: 110,
    treeRootFungi: 120,
    treeRootBacteria: 130,
    treeRootInsects: 140,
    treeRootAnimals: 150,
    treeRootWater: 160,
  };
  const trees: Tree[] = [
    {
      lat: 54.782792,
      lng: 9.424908,
      name: "Tree 0",
      id: 0,
      status: "healthy",
      neededWater: 100,
      ...treeMetaData,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
    },
    {
      lat: 54.782792,
      lng: 9.424908,
      name: "Tree 1",
      id: 1,
      status: "neutral",
      neededWater: 50,
      ...treeMetaData,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
    },
    {
      lat: 54.786913,
      lng: 9.408921,
      name: "Tree 2",
      id: 2,
      status: "unhealthy",
      neededWater: 0,

      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.78159,
      lng: 9.424873,
      name: "Tree 3",
      id: 3,
      status: "healthy",
      neededWater: 100,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.788817,
      lng: 9.425888,
      name: "Tree 4",
      id: 4,
      status: "neutral",
      neededWater: 50,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.782141,
      lng: 9.429827,
      name: "Tree 5",
      id: 5,
      status: "unhealthy",
      neededWater: 0,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.787557,
      lng: 9.438296,
      name: "Tree 6",
      id: 6,
      status: "healthy",
      neededWater: 100,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.811338,
      lng: 9.455175,
      name: "Tree 7",
      id: 7,
      status: "neutral",
      neededWater: 50,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.788306,
      lng: 9.44411,
      name: "Tree 8",
      id: 8,
      status: "unhealthy",
      neededWater: 0,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.805982,
      lng: 9.44757,
      name: "Tree 9",
      id: 9,
      status: "healthy",
      neededWater: 100,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.784699,
      lng: 9.438025,
      name: "Tree 10",
      id: 10,
      status: "neutral",
      neededWater: 50,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.760139,
      lng: 9.380937,
      name: "Tree 11",
      id: 11,
      status: "unhealthy",
      neededWater: 0,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.762046,
      lng: 9.385276,
      name: "Tree 12",
      id: 12,
      status: "healthy",
      neededWater: 100,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.761091,
      lng: 9.385719,
      name: "Tree 13",
      id: 13,
      status: "neutral",
      neededWater: 50,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.769905,
      lng: 9.47351,
      name: "Tree 14",
      id: 14,
      status: "unhealthy",
      neededWater: 0,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.771202,
      lng: 9.430948,
      name: "Tree 15",
      id: 15,
      status: "healthy",
      neededWater: 100,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.792941,
      lng: 9.462763,
      name: "Tree 16",
      id: 16,
      status: "neutral",
      neededWater: 50,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.797287,
      lng: 9.454632,
      name: "Tree 17",
      id: 17,
      status: "unhealthy",
      neededWater: 0,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
      ...treeMetaData,
    },
    {
      lat: 54.804966,
      lng: 9.486204,
      name: "Tree 18",
      id: 18,
      status: "healthy",
      neededWater: 100,
      ...treeMetaData,
      image: `/tree${Math.floor(Math.random() * (4 - 1 + 1) + 1)}.jpg`,
    },
  ];

  return (
    <TreeDataContext.Provider value={{ trees }}>
      {children}
    </TreeDataContext.Provider>
  );
};

export const useTreeDataContext = () => {
  const context = useContext(TreeDataContext);
  if (!context) {
    throw new Error(
      "useTreeDataContext must be used within a TreeDataContextProvider",
    );
  }
  return context;
};

export const useTrees = () => {
  const { trees } = useTreeDataContext();
  return trees;
};
