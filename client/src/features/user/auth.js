import { ApiUser } from "../../api/newapi"

const authApiSlice = ApiUser.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation({
      query: (body) => ({
        url: "/login",
        method: "POST",
        body: body
      })
    }),
    verify: builder.query({
      query: () => ({
        url: "/auth/verify",
        method: "GET"
      })
    }),
    sendEmail: builder.mutation({
      query: (body) => ({
        url: "/auth/verification/send",
        method: "POST",
        body: body
      })
    }),
    verifyOTP: builder.mutation({
      query: (body) => ({
        url: "/auth/verification/verify",
        method: "POST",
        body: body
      })
    }),
    setup: builder.mutation({
      query: (body) => ({
        url: "/auth/setup",
        method: "POST",
        body: body
      })
    }),
    register: builder.mutation({
      query: (body) => ({
        url: "/create",
        method: "POST",
        body: body
      })
    })
  })
})

export const {
  useLoginMutation,
  useRegisterMutation,
  useVerifyQuery,
  useSendEmailMutation,
  useVerifyOTPMutation,
  useSetupMutation
} = authApiSlice
