document.head.appendChild(
  Object.assign(document.createElement("style"), {
    type: "text/tailwindcss",
    textContent: `
        @layer components {
          .btn-internet {
            @apply outline outline-sky-400 p-1 bg-sky-200 rounded-md font-bold;
          }

          .btn-lg {
            @apply bg-sky-600 hover:bg-sky-700 rounded-md px-6 py-2
          }
        }`,
  })
);
