import {createLogin} from "./modules/createLogin.js";
import {createMainPage} from "./modules/createMainPage.js"
import {createRegistration} from "./modules/createRegistration.js";
import {createSettings} from "./modules/createSettings.js";
import {createProfile} from "./modules/createProfile.js";


const app = document.getElementById("application");

app.addEventListener('click', function (evt) {
  const {target} = evt;
  console.log(evt)
  if (target instanceof HTMLAnchorElement) {
    evt.preventDefault();
    routes[target.dataset.section](app);
  }
});

app.addEventListener('load', (event) => {
  const {target} = event;
  event.preventDefault();
});

const routes = {
  main: createMainPage,
  login: createLogin,
  registration: createRegistration,
  settings: createSettings,
  about: createProfile,
};

createMainPage(app);