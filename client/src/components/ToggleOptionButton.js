
import React, { useState } from 'react';

const ToggleOptionButton = ({ label }) => {
  const [isActive, setIsActive] = useState(false);

  return (
    <button
      onClick={() => setIsActive(!isActive)}
      className={`w-full flex items-center justify-between py-2 px-4 rounded-lg mt-2 ${isActive ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'
        }`}
    >
      <span>{label}</span>
    </button>
  );
};

export default ToggleOptionButton;
