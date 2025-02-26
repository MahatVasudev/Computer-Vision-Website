
import React from 'react';

const ImagePreview = () => {
  return (
    <div className="bg-gray-200 rounded-lg overflow-hidden h-full flex items-center justify-center">
      <img src="https://via.placeholder.com/300x500/8A2BE2/FFFFFF?text=Image+Preview" alt="Preview" className="object-cover" />
    </div>
  );
};

export default ImagePreview;
