import { faHome, faGear, faImage, faPlay } from "@fortawesome/free-solid-svg-icons";
import { wait } from "@testing-library/user-event/dist/utils";

const NavBarLeftConst = [
  {
    name: "Home",
    icon: faHome,
    href: "/",

  },
  {
    name: "Your Posts",
    icon: faImage,
    href: "/posts",

  },
  {
    name: "Recent",
    icon: faPlay,
    href: "/recent",

  },
  {
    name: "Settings",
    icon: faGear,
    href: "/setting",

  },

]

export {
  NavBarLeftConst
}
