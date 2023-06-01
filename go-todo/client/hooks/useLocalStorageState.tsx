import React, { useState, useRef, useEffect } from "react";

function useLocalStorageState<T>(
  key: string,
  defaultValue: T | (() => T) = "" as any,
  {
    serialize = JSON.stringify,
    deserialize = JSON.parse,
  }: {
    serialize?: (value: T) => string;
    deserialize?: (value: string) => T;
  } = {}
): [T, React.Dispatch<React.SetStateAction<T>>] {
  const [state, setState] = useState<T>(() => {
    const valueInLocalStorage = window.localStorage.getItem(key);
    if (valueInLocalStorage) {
      return deserialize(valueInLocalStorage);
    }
    return typeof defaultValue === "function"
      ? (defaultValue as () => T)()
      : defaultValue;
  });

  const prevKeyRef = useRef<string>(key);

  useEffect(() => {
    const prevKey = prevKeyRef.current;
    if (prevKey !== key) {
      window.localStorage.removeItem(prevKey);
    }
    prevKeyRef.current = key;
    window.localStorage.setItem(key, serialize(state));
  }, [key, state, serialize]);

  return [state, setState];
}

export { useLocalStorageState };
