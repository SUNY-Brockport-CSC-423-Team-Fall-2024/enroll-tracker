"use client";
import { useRouter } from "next/navigation";
import clsx from "clsx";
import { ButtonType } from "../lib/definitions";
import styles from "./styles.module.css";

interface IButton {
  btnTitle: string;
  href?: string;
  btnType: ButtonType;
}

export default function Button({ btnTitle, href, btnType }: IButton) {
  const router = useRouter();
  return (
    <button
      className={clsx(
        btnType === ButtonType.PRIMARY && styles.primary_button,
        btnType === ButtonType.SECONDARY && styles.secondary_button,
      )}
      onClick={() => {
        if (href) {
          router.push(href);
        }
      }}
    >
      {btnTitle}
    </button>
  );
}
