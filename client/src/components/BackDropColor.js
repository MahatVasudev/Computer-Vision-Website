import { useSelector } from "react-redux"

const BackDropColorBlur = () => {

  const preferedColor = useSelector((state) => state.settings.prefered_color)

  return (

    <div className={`fixed top-0 left-0 w-full h-[150px] bg-[radial-gradient(ellipse_at_top,rgba(255,255,255,0.3),transparent)] blur-3xl z-[-1] pointer-events-none`} style={{ backgroundColor: preferedColor }} />
  )
}

export default BackDropColorBlur
