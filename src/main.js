import {createMainPage} from "./modules/createMainPage";
import {createLogin} from "./modules/createLogin";
import {createRegistration} from "./modules/createRegistration";
import {createSettings} from "./modules/createSettings";
import  "./css/styles.css"
import  "./css/normalize.css"
import {createProfile} from "./modules/createProfile";


const app = document.getElementById("application");

app.addEventListener('click', function (evt) {
  const {target} = evt;
  if (target instanceof HTMLAnchorElement) {
    evt.preventDefault();
    routes[target.getAttribute('section')](app);
  }
});

app.addEventListener('load', (event) => {
  const {target} = event;
  console.log(target)
  event.preventDefault();
});

const routes = {
  main: createMainPage,
  login: createLogin,
  registration: createRegistration,
  settings: createSettings,
  // about: createProfile,
};

createMainPage(app);
