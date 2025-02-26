import React, { useState } from 'react';
import ToggleOptionButton from '../components/ToggleOptionButton';
import ImageEnhancerButton from '../components/ImageEnhancerButton';
import { faAdjust, faSlidersH, faSun, faTint } from '@fortawesome/free-solid-svg-icons';

const EditImage = () => {
  const Image_Enhancer = [
    { title: 'Brightness', icon: faSun },
    { title: 'Contrast', icon: faAdjust },
    { title: 'Saturation', icon: faTint },
    { title: 'Sharpness', icon: faSlidersH }
  ];

  const [selectedEnhancer, setSelectedEnhancer] = useState(Image_Enhancer[0].title);

  const handleEnhancerClick = (title) => {
    setSelectedEnhancer(title); // Set the selected enhancer to the clicked one
  };

  return (
    <div className="flex h-screen my-4 mx-5">
      {/* Image Section */}
      <img
        src="https://via.placeholder.com/300x500/8A2BE2/FFFFFF?text=Image"
        alt="Editable Post"
        className="w-1/2 h-full object-cover rounded-md"
      />

      {/* Edit Post Section */}
      <div className="w-1/2 ml-8 bg-gray-50 shadow-lg rounded-lg p-6 flex flex-col relative">
        <h2 className="text-xl font-semibold mb-4">Edit Post</h2>

        {/* Scrollable Content */}
        <div className="overflow-y-auto flex-grow pr-2">
          <ToggleOptionButton label="Auto Colorization" />
          <ToggleOptionButton label="Auto Background Blur" />

          <div className="mt-4">
            <label className="block text-gray-700">{selectedEnhancer}</label>
            <input type="range" className="w-full mt-2" />
          </div>

          <div className="mt-4">
            <h3 className="text-gray-700 font-medium">Image Enhancers</h3>
            <div className="grid grid-cols-4 gap-4 mt-2">
              {Image_Enhancer.map((enhancer, key) => (
                <ImageEnhancerButton
                  key={key}
                  label={enhancer.title}
                  icon={enhancer.icon}
                  isSelected={selectedEnhancer === enhancer.title}
                  onClick={() => handleEnhancerClick(enhancer.title)}
                />
              ))}
            </div>

          </div>
          {/* Fixed Post Button */}
        </div>
        <button className="bg-blue-500 text-white rounded-full py-2 mt-auto">
          Post
        </button>



      </div>
    </div>
  );
};

export default EditImage;
