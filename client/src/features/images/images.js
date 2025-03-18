
import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  filters: {
    brightness: 100,
    contrast: 100,
    saturation: 100,
    sharpness: 0,
    grain: 0,
    highlights: 0,
    shadows: 0,
    dispersion: 0,
  },
};

const imageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    updateFilter: (state, action) => {
      state.filters = action.payload;
    },
  },
});

export const { updateFilter } = imageSlice.actions;
export default imageSlice.reducer;
