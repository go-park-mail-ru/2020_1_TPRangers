export function createPost( parent = document.body,
  data = {
  name: 'Default post name',
  textData: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.',
  imageData: 'https://picsum.photos/200/300?grayscale',
}) {

  const postCard = document.createElement('div');
  postCard.classList.add('postCard');

  const image = document.createElement('img');
  image.classList.add('image');
  image.src = data.imageData;

  const name = document.createElement('span');
  name.classList.add('postName');
  name.textContent = data.name;

  const text = document.createElement('p');
  text.classList.add('postText');
  text.textContent = data.textData;

  postCard.appendChild(name);
  postCard.appendChild(image);
  postCard.appendChild(text);

  parent.appendChild(postCard);
}