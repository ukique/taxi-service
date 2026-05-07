import { useContext } from 'react';
import { WSContext } from './WSContext';

export const useWS = () => useContext(WSContext);