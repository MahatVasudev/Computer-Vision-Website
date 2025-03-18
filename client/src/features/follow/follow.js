import { ApiFollow } from "../../api/newapi";

const FollowApiSlice = ApiFollow.injectEndpoints({
  endpoints: builder => ({
    followUser: builder.mutation({
      query: (body) => ({
        url: '/follow',
        method: 'POST',
        body: body
      })
    })
  })
})


export const { useFollowUserMutation } = FollowApiSlice

