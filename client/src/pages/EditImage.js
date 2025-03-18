import { useLocation, useNavigate } from "react-router-dom";
import { useRef, useState, useEffect } from "react";
import Button from "../components/button";

const filterOptions = [
  { name: "Brightness", key: "brightness", min: 0, max: 200, default: 100 },
  { name: "Contrast", key: "contrast", min: 0, max: 200, default: 100 },
  { name: "Saturation", key: "saturation", min: 0, max: 200, default: 100 },
];

const EditImage = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const canvasRef = useRef(null);
  const ctxRef = useRef(null);
  const originalImage = useRef(null); // Stores original image for resets
  const [selectedFilter, setSelectedFilter] = useState(null);
  const [filterValues, setFilterValues] = useState({
    brightness: 100,
    contrast: 100,
    saturation: 100,
  });
  const { imageSrc } = location.state || {};

  useEffect(() => {
    if (!imageSrc) return;
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");
    ctxRef.current = ctx;

    const image = new Image();
    image.src = imageSrc;
    image.onload = () => {
      canvas.width = image.width;
      canvas.height = image.height;
      ctx.drawImage(image, 0, 0);
      originalImage.current = image; // Save original image
    };
  }, [imageSrc]);

  const applyFilters = () => {
    const canvas = canvasRef.current;
    const ctx = ctxRef.current;
    if (!canvas || !ctx || !originalImage.current) return;

    // Clear and reset to original image
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.drawImage(originalImage.current, 0, 0);

    // Apply filters
    ctx.filter = `
      brightness(${filterValues.brightness}%)
      contrast(${filterValues.contrast}%)
      saturate(${filterValues.saturation}%)
    `;
    ctx.drawImage(originalImage.current, 0, 0);
  };

  useEffect(() => {
    applyFilters();
  }, [filterValues]);

  const resetFilters = () => {
    setFilterValues({
      brightness: 100,
      contrast: 100,
      saturation: 100,
    });
  };

  const saveEditedImage = () => {
    const canvas = canvasRef.current;

    canvas.toBlob((blob) => {
      const file = new File([blob], "edited-image.png", { type: "image/png" });

      navigate("/newpost", { state: { editedImageFile: file } });
    }, "image/png");
    // Navigate back with edited image URL

  };

  const AutoColorImage = () => {


    const canvas = canvasRef.current;

    canvas.toBlob((blob) => {
      const file = new File([blob], "edited-image.png", { type: "image/png" });

      navigate("/newpost", { state: { editedImageFile: file } });
    }, "image/png");

    const formdata = new FormData()
    formdata.append("file", file)

    try {

    } catch (e) {

    }
  }

  return (
    <div className="flex flex-col items-center bg-gray-100 p-6 w-full h-screen">
      {/* Top Slider */}
      {selectedFilter && (
        <div className="w-full max-w-2xl mb-4 px-4">
          <input
            type="range"
            min={filterOptions.find((f) => f.key === selectedFilter).min}
            max={filterOptions.find((f) => f.key === selectedFilter).max}
            value={filterValues[selectedFilter]}
            onChange={(e) =>
              setFilterValues({ ...filterValues, [selectedFilter]: Number(e.target.value) })
            }
            className="w-full appearance-none h-2 bg-gray-300 rounded-lg cursor-pointer accent-blue-500"
          />
        </div>
      )}

      {/* Image Canvas */}
      <div className="relative flex justify-center items-center w-[500px] h-[500px] bg-white shadow-md border">
        <canvas ref={canvasRef} className="max-w-full max-h-full"></canvas>
      </div>

      {/* Filter Selection */}
      <div className="flex justify-center gap-4 mt-6">
        {filterOptions.map((filter) => (
          <button
            key={filter.key}
            className={`px-6 py-2 rounded-md text-sm font-medium transition-all duration-200 ${selectedFilter === filter.key
              ? "bg-blue-500 text-white shadow-md"
              : "bg-gray-200 text-gray-800 hover:bg-gray-300"
              }`}
            onClick={() => setSelectedFilter(filter.key)}
          >
            {filter.name}
          </button>
        ))}
        <button className={`px-6 py-2 rounded-md text-sm font-medium transition-all duration-200`} onClick={() => }></button>
      </div>

      <div className="flex gap-4 mt-6">
        <button onClick={resetFilters} className="px-4 py-2 bg-gray-500 text-white rounded-md">
          Reset
        </button>
        <button onClick={saveEditedImage} className="px-4 py-2 bg-blue-500 text-white rounded-md">
          Save
        </button>
      </div>
    </div>
  );
};

export default EditImage;
