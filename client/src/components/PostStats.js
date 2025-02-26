import React from 'react';
import { faHeart, faComment, faBookmark, faDownload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const PostStats = ({ isPostLiked, isPostBookmarked, toggleLike, toggleBookmark, imageUrl }) => {

  // Function to handle image download using fetch and blob
  const downloadImage = async (event) => {
    event.preventDefault(); // Prevent default anchor behavior

    try {
      // Check if the image is from the same origin or if CORS headers allow fetching
      const response = await fetch(imageUrl, {
        mode: 'cors', // Ensure we handle CORS
        headers: {
          'Access-Control-Allow-Origin': '*',
        }
      });

      // Check if the response is OK (status 200-299) and handle blob creation
      if (response.ok) {
        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);

        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'downloaded_image.jpg'); // Set the download attribute
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link); // Clean up by removing the link element
        window.URL.revokeObjectURL(url); // Clean up the object URL
      } else {
        console.error('Failed to fetch image:', response.statusText);
        // Fallback: Open the image in a new tab if fetch fails
        window.open(imageUrl, '_blank');
      }
    } catch (error) {
      console.error('Error downloading the image:', error);
      // Fallback: Open the image in a new tab in case of a CORS error or fetch failure
      window.open(imageUrl, '_blank');
    }
  };

  return (
    <div className="border-t py-4">
      <div className="flex items-center justify-between">
        <div className="flex space-x-4">
          <FontAwesomeIcon
            icon={faHeart}
            className={`cursor-pointer ${isPostLiked ? 'text-red-500' : 'text-gray-600'}`}
            onClick={toggleLike}
          />
          <FontAwesomeIcon icon={faComment} className="text-gray-600" />
        </div>
        <div className="ml-auto flex space-x-4">
          {/* Download button with image download functionality */}
          <FontAwesomeIcon
            icon={faDownload}
            className="cursor-pointer text-gray-600"
            onClick={downloadImage} // Trigger the download when clicked
          />
          <FontAwesomeIcon
            icon={faBookmark}
            className={`cursor-pointer ${isPostBookmarked ? 'text-yellow-500' : 'text-gray-600'}`}
            onClick={toggleBookmark}
          />
        </div>
      </div>
      <p className="text-sm text-gray-500 mt-2">10 Likes · 3 Comments · 2 Saves</p>
    </div>
  );
};

export default PostStats;
