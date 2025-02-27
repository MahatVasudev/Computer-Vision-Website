
const authMiddleware = (store) => (next) => (action) => {
  if (action.type.endsWith("/rejected")) {
    const status_ = action.payload?.response?.status;

    if (status_ === 401) {
      window.location.href = "/login"
    }
  }

  return next(action)
}

export default authMiddleware
