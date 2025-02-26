import { ApiUser } from "../../api/newapi";

const ApiUserEndPoints = ApiUser.injectEndpoints({
  endpoints: builder => ({
    userdetails: builder.mutation({
      query: ({ username }) => ({
        url: `/details/un/${username}`,
        method: "GET"
      })
    })
  })
})

export const { useUserdetailsMutation } = ApiUserEndPoints
