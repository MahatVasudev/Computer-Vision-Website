import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

const ImageEnhancerButton = ({ label, icon, isSelected, onClick }) => {
  return (
    <button
      onClick={onClick}
      className={`flex flex-col items-center justify-center w-10 h-10 rounded-lg ${isSelected ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-700'
        }`}
    >
      <FontAwesomeIcon icon={icon} className="mb-1" />
      <span className="text-xs">{label}</span>
    </button>
  );
};

export default ImageEnhancerButton;
