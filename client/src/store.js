import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { FLUSH, PAUSE, PERSIST, PURGE, REGISTER, REHYDRATE, persistStore } from "redux-persist";
import persistReducer from "redux-persist/es/persistReducer";
import storage from "redux-persist/lib/storage";
import { ApiPost, ApiUser } from "./api/newapi";
import settingsLocal from "./features/settings/settings.local";
import authLocal from "./features/user/auth.local";
import authMiddleware from "./middlewares/auth.middleware";

const persistConfig = {
  key: "root",
  version: "1",
  storage: storage,
  blaclist: [],
}

const reducers = combineReducers({
  "auth": authLocal,
  "settings": settingsLocal,
  [ApiUser.reducerPath]: ApiUser.reducer,
  [ApiPost.reducerPath]: ApiPost.reducer
})

const persistedReducer = persistReducer(persistConfig, reducers)

const store = configureStore({
  reducer: persistedReducer,
  middleware: getDefaultMiddleWare => getDefaultMiddleWare({
    serializableCheck: {
      ignoreActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER]
    }
  }).concat([ApiUser.middleware, ApiPost.middleware, authMiddleware])
})


export default store

export const persistedStore = persistStore(store)
