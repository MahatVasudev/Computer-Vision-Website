
import React from 'react';

import Masonry from "react-masonry-css";
import { getPosts } from '../utils/common';


const MasonryGrid = ({ data }) => {
  const breakpointColumnsObj = {
    default: 6,  // 4 columns by default
    1800: 5,
    1300: 4,
    1100: 3,     // 3 columns for screens <= 1100px
    700: 2,      // 2 columns for screens <= 700px
  };

  if (data == null || data.length == 0) {
    return (
      <h1 className='dark:text-white text-black'>No Posts Made yet!!</h1>
    )
  }

  return (

    <Masonry
      breakpointCols={breakpointColumnsObj}
      className="my-masonry-grid"
      columnClassName="my-masonry-grid_column"
    >
      {data.map((d, index) => (
        <div key={index} className="mb-4">
          <a href={`/post/${d.post_id}`}>
            <img src={getPosts(d.image_location)} alt={`Image ${index + 1}`} className="w-full h-auto rounded-lg shadow-md" />
          </a>
        </div>
      ))}
    </Masonry>
  );
};

export default MasonryGrid;
