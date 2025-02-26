import { createSlice } from "@reduxjs/toolkit";

const localSettings = createSlice({
  name: "settings",
  initialState: {
    theme: "light",
  },
  reducers: {
    setSettings: (state, action) => {
      const { key, value } = action.payload

      state[key] = value
    },
    getSettings: (state, action) => {
      const { key } = action.payload

      return state[key]
    },
    delSettings: (state, action) => {
      const { key } = action.payload

      delete state[key]
    }
  }

})

export const { setSettings } = localSettings.actions

export default localSettings.reducer
