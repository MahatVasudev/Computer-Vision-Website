import { createSlice } from "@reduxjs/toolkit";

const localSettings = createSlice({
  name: "settings",
  initialState: {
    dark: 0,
    prefered_color: "#000000"
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
    },
    toggleThemes: (state, action) => {
      state.dark = state.dark === 1 ? 0 : 1
    }
  }

})

export const { setSettings, toggleThemes } = localSettings.actions

export default localSettings.reducer
