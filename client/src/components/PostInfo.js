
import React, { useState } from 'react';

const PostInfo = ({ creator, title, shortDescription, fullDescription }) => {
  const [isDescriptionExpanded, setIsDescriptionExpanded] = useState(false);

  const toggleDescription = () => {
    setIsDescriptionExpanded(!isDescriptionExpanded);
  };

  const shouldShowMorePostDesc = fullDescription.length > shortDescription.length;

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{title}</h1>
      <div className="flex items-center mb-4">
        <img
          src={creator.avatarUrl || '/user_default.jpg'} // Default avatar if not provided
          alt="Creator Avatar"
          className="h-10 w-10 rounded-full border"
        />
        <div className="ml-2">
          <p className="font-semibold">@{creator}</p>
          <p className="text-sm text-gray-500">
            {isDescriptionExpanded ? fullDescription : `${shortDescription} `}
            {shouldShowMorePostDesc && (
              <span className="text-blue-500 cursor-pointer" onClick={toggleDescription}>
                {isDescriptionExpanded ? "Show less" : "...more"}
              </span>
            )}
          </p>
        </div>
      </div>
    </div>
  );
};

export default PostInfo;
