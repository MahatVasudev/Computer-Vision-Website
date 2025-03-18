import { ApiUser } from "../../api/newapi";

const ApiUserEndPoints = ApiUser.injectEndpoints({
  endpoints: builder => ({
    userdetails: builder.mutation({
      query: ({ username }) => ({
        url: `/details/un/${username}`,
        method: "GET"
      })
    }),

    GetAllUserPosts: builder.mutation({
      query: (username) => ({
        url: `/details/un/${username}/posts`,
        method: "GET"
      })
    }),
    seeUsernameExists: builder.mutation({
      query: ({ username }) => ({
        url: '/check/un',
        method: "GET",
        body: { username }
      })
    })
  })
})


export const { useGetAllUserPostsMutation, useUserdetailsMutation, useSeeUsernameExistsMutation } = ApiUserEndPoints
