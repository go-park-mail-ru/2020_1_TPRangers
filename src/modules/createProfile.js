import {createPost} from './createPost.js';
import createLinks from "./createLinks.js";

export function createProfile(parent, user = {
  name: 'UserName',
  dateOfB: '00',
  monthOfB: '00',
  yearOfB: '0000',
  avatar: 'https://picsum.photos/200/300'
}) {
  parent.innerHTML = '';

  const leftBlock = document.createElement('div');
  leftBlock.classList.add('leftBlock');

  const rightBlock = document.createElement('div');
  rightBlock.classList.add('rightBlock');

  const avatar = document.createElement('img');
  avatar.classList.add('userAvatar');
  avatar.src = user.avatar;

  const name = document.createElement('span');
  name.classList.add('userName');
  name.textContent = user.name;

  const dateOfBLabel = document.createElement('span');
  dateOfBLabel.classList.add('dateOfBLabel');
  dateOfBLabel.textContent = 'Date of birth: ';

  const dateOfB = document.createElement('span');
  dateOfB.classList.add('dateOfBUser');
  dateOfB.textContent = `${user.dateOfB}.${user.monthOfB}.${user.yearOfB}`;

  rightBlock.appendChild(name);
  rightBlock.appendChild(dateOfBLabel);
  rightBlock.appendChild(dateOfB);
  rightBlock.appendChild(createLinks({
    name: 'Редактировать профиль',
    link: 'settings',
    cl: 'userSettings'
  }));
  for (let i = 0; i < 10; ++i) {
    createPost(rightBlock);
  }

  leftBlock.appendChild(avatar);

  parent.appendChild(leftBlock);
  parent.appendChild(rightBlock);
}