import api from "../../common/axios";

export const logIn = async (data: any) => {
  return await api.post("log-in", data);
};

