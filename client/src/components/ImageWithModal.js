
import React, { useState } from 'react';

const ImageWithModal = ({ imageUrl }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  console.log("Image Url", imageUrl)
  return (
    <>
      <div className="w-fit mx-auto bg-transparent h-fit py-5 shadow-lg rounded-lg flex justify-center items-center">
        <img
          src={imageUrl}
          alt="Post"
          className="h-[40rem] rounded-lg cursor-pointer"
          onClick={openModal} // Open modal when image is clicked
        />
      </div>

      {/* Full-Screen Image Modal */}
      {isModalOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-70 flex justify-center items-center z-50"
          onClick={closeModal} // Close modal when background is clicked
        >
          <img
            src={imageUrl}
            alt="Full-screen Post"
            className="w-full h-full object-contain"
          />
        </div>
      )}
    </>
  );
};

export default ImageWithModal;
