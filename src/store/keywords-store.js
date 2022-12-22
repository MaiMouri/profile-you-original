import { configureStore } from '@reduxjs/toolkit';
import keywordsReducer from '../keywordsSlice';


export default configureStore({
    reducer: {
      keywords: keywordsReducer,
    },
  });

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