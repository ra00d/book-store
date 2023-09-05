import { ReactNode } from "react";
import { RouteObject } from "react-router-dom";

export interface IRoute {
  title: string;
  icon: ReactNode;
  path: string;
}

export type TRoute = RouteObject & {
  title: string;
  icon: ReactNode;
  path: string;
};
