import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import Button from "../components/button";
import { Pen } from "lucide-react";
import Input from "../components/InputTextFields";
import TextAreaField from "../components/TextArea";
import { useCreatePostMutation } from "../features/post/posts";

export default function CreatePost() {
  const location = useLocation();
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const navigate = useNavigate();
  const [editedImage, setEditedImage] = useState();
  const [editedFile, setEditedFile] = useState()
  const [createpost, { isLoading }] = useCreatePostMutation()



  useEffect(() => {
    if (location.state?.editedImageFile) {
      setEditedFile(location.state?.editedImageFile)
      setEditedImage(URL.createObjectURL(location.state.editedImageFile));
    }
  }, [location.state])


  const PostPost = async () => {

    console.log("clicked!")

    try {

      const formdata = new FormData()

      formdata.append("title", title)
      formdata.append("description", description)
      formdata.append("images", editedFile)

      const res = await createpost(formdata)

      if (res.error) {
        console.log(res.error)
        alert(`Error Occured: ${res.error.data.error}, please try again`)
        return
      }

      const post_id = res.data.data.post_id
      console.log(post_id)
      navigate(`/post/${post_id}`)

    } catch (e) {
      console.log(e)
    }


  }

  const handleDrop = (event) => {
    event.preventDefault();
    const file = event.dataTransfer.files[0];
    if (file) {
      setEditedImage(URL.createObjectURL(file));
      setEditedFile(file);
    }
  };

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      setEditedImage(URL.createObjectURL(file));
      setEditedFile(file);
    }
  };

  return (
    <div className="flex flex-col md:flex-row gap-8 p-6">
      {/* Image Upload Section */}
      <div className="flex flex-col items-center gap-4 p-4 border-2 border-dashed border-gray-300 rounded-lg w-full md:w-1/2"
        onDrop={handleDrop}
        onDragOver={(event) => event.preventDefault()}
      >
        {editedImage ? (
          <img src={editedImage} alt="Edited" className="mt-auto mb-auto rounded-md" />) :
          (
            <div className="text-center text-gray-500 mt-auto mb-auto">Drag & Drop or Click to Upload</div>
          )}
        <label className="cursor-pointer">
          <input type="file" className="hidden" onChange={handleFileChange} accept="image/*" />
        </label>
      </div>

      {/* Form Section */}
      <div className="flex flex-col h-[80vh] w-full md:w-1/2 p-4 bg-gray-100 rounded-lg">
        <h2 className="text-xl font-semibold mb-4">Creating Post</h2>
        <h3 className="text-lg font-semibold mb-3">Title</h3>
        <Input state={{ state: title, setState: setTitle }} />

        <h3 className="text-lg font-semibold mb-3">Description</h3>
        <TextAreaField state={{ state: description, setState: setDescription }} />

        <div className="flex flex-row mt-auto justify-between">
          {/* Edit Image Button */}
          <Button
            variant="outline"
            color="white"
            text={<><Pen /></>}
            onClick={() =>
              navigate("/edit", { state: { imageSrc: editedImage, title, description } })
            }
          />
          <Button text="Post" color="white" onClick={
            () => { PostPost() }

          } />
        </div>
      </div>
    </div>
  );
}
