import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react"

const local = "http://localhost:5000"

const baseQueryUser = fetchBaseQuery({
  baseUrl: local + "/user",
  credentials: 'include',
})

const baseQueryPost = fetchBaseQuery({
  baseUrl: local + "/post"
})

const ApiUser = createApi({
  reducerPath: "user_api",
  baseQuery: baseQueryUser,
  endpoints: () => ({}),
})

const ApiPost = createApi({
  reducerPath: "post_api",
  baseQuery: baseQueryPost,
  endpoints: () => ({})
})

export { ApiUser, ApiPost }
