import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import { routes } from "../routes/routes";
// import { TRoute } from "../../types/routes";

export const routesStore = create<any>()(
  devtools(
    persist(
      () => ({
        // routes: routes,
        getRoutes: () => {
          return routes;
        },
      }),
      { name: "settings" },
    ),
  ),
);
