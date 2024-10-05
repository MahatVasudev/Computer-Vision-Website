import React from "react";
import MasonryGrid from "../components/Masonry_Grid";

const Home = () => {
  const images = [
    "https://via.placeholder.com/300x400/FF7F50/FFFFFF?text=Image+1",
    "https://via.placeholder.com/300x300/6495ED/FFFFFF?text=Image+2",
    "https://via.placeholder.com/300x500/FF69B4/FFFFFF?text=Image+3",
    "https://via.placeholder.com/300x350/8A2BE2/FFFFFF?text=Image+4",
    "https://via.placeholder.com/300x450/00FA9A/FFFFFF?text=Image+5",
    "https://via.placeholder.com/300x300/FFD700/FFFFFF?text=Image+6",
    "https://via.placeholder.com/300x400/DC143C/FFFFFF?text=Image+7",
    "https://via.placeholder.com/300x350/ADFF2F/FFFFFF?text=Image+8",
    "https://via.placeholder.com/300x500/20B2AA/FFFFFF?text=Image+9",
    "https://via.placeholder.com/300x400/FF7F50/FFFFFF?text=Image+10",
  ];

  // Define breakpoint columns for 4 columns layout

  return (
    <div className="mx-[8rem]">
      <MasonryGrid data={images} />
    </div>

  );
};

export default Home;
