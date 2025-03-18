import { setCreds } from "../features/user/auth.local"
const authMiddleware = (store) => (next) => (action) => {


  if (action.type.endsWith("/rejected")) {

    const status_ = action?.meta?.baseQueryMeta?.response?.status;

    if (status_ === 401) {

      window.location.href = "/login"
      store.dispatch(setCreds({ key: "status", value: false }))
      return
    }
  }

  return next(action)
}

export default authMiddleware
