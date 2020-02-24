export default function createLinks(data = {
  by_default: {
    name: 'undefined',
    link: 'undefined',
    cl: 'main_link'

  } }) {
  const pageItem = document.createElement('a');
  pageItem.textContent = data.name;
  pageItem.href = `/${data.link}`;
  pageItem.dataset.section = data.link;
  pageItem.classList.add(data.cl);

  return pageItem;
}