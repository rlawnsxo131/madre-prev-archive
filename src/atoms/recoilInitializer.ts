import { MutableSnapshot } from 'recoil';

export default function recoilInitializer({ set }: MutableSnapshot) {
  console.log(set);
}
