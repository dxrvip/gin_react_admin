import { AuthProvider, HttpError } from "react-admin";


interface LoginResponse {
  status: number;
  msg: string;
  data: { user_id: number, username: string, full_name: string };
  token: string;
}
/**
 * 此 authProvider 仅用于测试目的。不要在生产中使用它。
 */
export const authProvider: AuthProvider = {
  login: ({ username, password }) => {
    const request = new Request(import.meta.env.VITE_SIMPLE_REST_URL + "/user/login", {
      method: "POST",
      body: JSON.stringify({ username, password }),
      headers: new Headers({ "Content-Type": "application/json", "accept": "application/json" }),
    });

    let msg: string = "服务器错误";
    return fetch(request)
      .then(async (response) => {
        // 访问成功
        if (response.status === 200) {
          // eslint-disable-next-line no-unused-vars, @typescript-eslint/no-unused-vars
          let result: LoginResponse = await response.json();
          if (result.status === 200) {
            // 登陆成功//存储token
            localStorage.setItem("user", JSON.stringify(result?.data));
            localStorage.setItem("token", result?.token);
            return Promise.resolve();

          }
          msg = result.msg;
        }
        return Promise.reject();
      }).catch(() => {
        return Promise.reject(
          new HttpError(msg, 500, {
            message: msg,
          })
        );
      })



  },
  logout: () => {
    localStorage.removeItem("user");
    localStorage.removeItem("token");
    return Promise.resolve();
  },

  checkError: (error) => {
    console.log("========>", error);
    const status = error.status;
    if (status === 401 || status === 403) {
      localStorage.removeItem("user");
      localStorage.removeItem("token");
      return Promise.reject();
    }

    return Promise.resolve()
  },
  checkAuth: () => {
    // 判断是否有token和用户
    let verify = (localStorage.getItem("token") == null || localStorage.getItem("user") == null)
    if (verify) {
      return Promise.reject()
    }
    return Promise.resolve()
  },
  getPermissions: () => {
    return Promise.resolve(undefined);
  },
  getIdentity: () => {
    const persistedUser = localStorage.getItem("user");
    const user = persistedUser ? JSON.parse(persistedUser) : null;

    return Promise.resolve(user);
  },
};

export default authProvider;