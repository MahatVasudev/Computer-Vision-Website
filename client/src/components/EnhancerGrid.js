
import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSun, faAdjust, faTint, faSlidersH } from '@fortawesome/free-solid-svg-icons';

const enhancers = [
  { icon: faSun, label: "Brightness" },
  { icon: faAdjust, label: "Contrast" },
  { icon: faTint, label: "Saturation" },
  { icon: faSlidersH, label: "Sharpness" },
  // Add other enhancers here
];

const EnhancerGrid = () => {
  return (
    <div className="grid grid-cols-4 gap-2 mb-4">
      {enhancers.map((enhancer, index) => (
        <button key={index} className="p-2 bg-gray-200 rounded-lg flex flex-col items-center">
          <FontAwesomeIcon icon={enhancer.icon} className="text-gray-700 mb-1" />
          <span className="text-xs">{enhancer.label}</span>
        </button>
      ))}
    </div>
  );
};

export default EnhancerGrid;
