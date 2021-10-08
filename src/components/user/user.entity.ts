import {
  Column,
  CreateDateColumn,
  Entity,
  Index,
  OneToMany,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import { Data } from '../data';

@Entity('user')
class User {
  @PrimaryGeneratedColumn({ unsigned: true })
  id!: number;

  @Index('ix_email', { unique: true })
  @Column()
  email!: string;

  @Column({ length: 16, nullable: true, default: null })
  username?: string;

  @Column({ length: 48 })
  display_name!: string;

  @Column({ length: 255, nullable: true, default: null })
  photo_url?: string;

  @CreateDateColumn({ type: 'timestamp' })
  created_at!: Date;

  @UpdateDateColumn({ type: 'timestamp' })
  updated_at!: Date;

  @OneToMany(() => Data, (data) => data.user)
  datas?: Data[];
}

export default User;
