
import React from 'react';

const Slider = ({ value, onValueChange, min, max }) => {
  return (
    <input
      type="range"
      className="w-full"
      min={min}
      max={max}
      value={value}
      onChange={onValueChange}
    />
  );
};

export default Slider;
