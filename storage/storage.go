/*
 * This file is part of Crocodile Game Bot.
 * Copyright (C) 2019  Viktor
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package storage

import "github.com/gomodule/redigo/redis"

type Storage struct {
	*Postgres
	*Redis
}

func NewStorage(conn string, pool *redis.Pool, logger Logger) (*Storage, error) {
	pg, err := NewPostgres(conn, logger)
	if err != nil {
		return &Storage{}, err
	}

	redis := NewRedis(pool)
	return &Storage{
		Postgres: pg,
		Redis:    redis,
	}, nil
}
