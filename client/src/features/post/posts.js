import { ApiPost } from "../../api/newapi";

const ApiPostInject = ApiPost.injectEndpoints({
  endpoints: (builder) => ({
    CreatePost: builder.mutation({
      query: (body) => ({
        url: '/make',
        method: "POST",
        body: body,
        formData: true
      })
    }),
    GetPostDetails: builder.mutation({
      query: (post_id) => ({
        url: `/i/${post_id}`,
        method: "GET"
      })
    }),
    GetAllPosts: builder.mutation({
      query: () => ({
        url: '/all',
        method: "GET"
      })
    })
  })
})


export const { useCreatePostMutation, useGetAllPostsMutation, useGetPostDetailsMutation } = ApiPostInject
