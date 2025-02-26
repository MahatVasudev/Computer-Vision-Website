import React from "react";
import MasonryGrid from "../components/Masonry_Grid";

const Home = () => {

  const images = [
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRRCZVKWKAUmqHUszu8_M3CoepdRNIXk9SvZQ&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSNxnwEFPxa4ujQwSb3ebJQ0qScgFz7CEY14g&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTrQSQC8BQOCvipNCu7-MdwTbVjqfASDE0pA&s",
    "https://wallpapers.com/images/hd/pink-aesthetic-anime-phone-zrwdo7l6d7dshf6k.jpg",
    "https://mfiles.alphacoders.com/101/thumb-1920-1013619.png",
    "https://lh3.googleusercontent.com/ZX6QtuP7xer0M5-ov_vw7K4qgjER0j_CnV5XoQ2KZU6DD2F4eRg2OyCeknEjvO-Uoyjk7_x5ljPJQe1e7F85aAdgMQ=s1280-w1280-h800",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS3ddrAtZQNd84Q2Ka_HBFozDlf81m2N7bEHg&s",
    "https://images7.alphacoders.com/135/thumb-1920-1352330.jpeg",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCMQrAsyMardvp4iH8tqoypTJJHtZ9e5jfVg&s",
    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRfwEyltdnZT7_9BDyPtgUW8aXhdq5R3AkRqw&s",
  ];
  // Define breakpoint columns for 4 columns layout

  return (
    <div className="mx-[8rem]">
      <MasonryGrid data={images} />
    </div>

  );
};

export default Home;
