import { configureStore } from '@reduxjs/toolkit';
import keywordsReducer from '../keywordsSlice';
import userReducer from '../redux/feature/user/userSlice'


export default configureStore({
    reducer: {
      keywords: keywordsReducer,
      user: userReducer,
    },
  });

