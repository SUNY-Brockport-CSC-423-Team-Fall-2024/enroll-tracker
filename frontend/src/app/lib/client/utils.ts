import { Dispatch, SetStateAction } from "react";

const handleTextInputChange = (
  reactHook: Dispatch<SetStateAction<any>>,
  e: React.FormEvent<HTMLInputElement | HTMLTextAreaElement>,
  isNumber?: boolean,
) => {
  reactHook(isNumber ? Number(e.currentTarget.value) : e.currentTarget.value);
};

export { handleTextInputChange };
