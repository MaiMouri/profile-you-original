// import { initStore } from "./store";

// const configureStore = () => {
//     const actions = {
//         TOGGLE_FAV: (curState, keywordId) => {
//             const keyIndex = curState.keywords.findIndex(k =>k.id === keywordId)
//             const newFavStatus = !curState.keywords[keyIndex].isFavorite;
//             const updatedKeywords = [...curState.keywords]
//             updatedKeywords[keyIndex] = {
//                 ...curState.keywords[keyIndex],
//                 isFavorite: newFavStatus
//             }
//             return { keywords: updatedKeywords}
//         }
//     }
//     initStore();
// }