import { createSlice } from "@reduxjs/toolkit";

const localAuthSlice = createSlice({
  name: 'auth',
  initialState: {
    status: false
  },
  reducers: {
    setCreds: (state, action) => {
      const { key, value } = action.payload

      state[key] = value
    },

    getCreds: (state, action) => {
      const { key } = action.payload

      return state[key]
    },
    delCreds: (state, action) => {
      const { key } = action.payload

      delete state[key]
    }
  }
})

export const { getCreds, setCreds, delCreds } = localAuthSlice.actions

export default localAuthSlice.reducer
