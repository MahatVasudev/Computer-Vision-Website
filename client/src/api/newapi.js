import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react"

export const local = "http://localhost:5000"

export const local_public = local + "/public"

export const local_post = local_public + "/posts"


const baseQueryUser = fetchBaseQuery({
  baseUrl: local + "/user",
  credentials: 'include',
})

const baseQueryPost = fetchBaseQuery({
  baseUrl: local + "/posts",
  credentials: 'include'
})

const baseQueryFollow = fetchBaseQuery({
  baseUrl: local + "/follow",
  credentials: 'include'
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


const ApiFollow = createApi({
  reducerPath: "follow_api",
  baseQuery: baseQueryFollow,
  endpoints: () => ({})
})

export { ApiUser, ApiPost, ApiFollow }
